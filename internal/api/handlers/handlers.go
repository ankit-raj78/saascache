package handlers

import (
    "context"
    "math/rand"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/yourorg/saascache/internal/provision"
)

// Health returns 200 OK for liveness checks.
func Health(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Recommend is a stub for the fast, request‑driven recommendation.
func Recommend(c *gin.Context) {
    // TODO: bind input JSON & call recommendation engine
    c.JSON(http.StatusOK, gin.H{
        "engine":  "in-memory",
        "pattern": "cache-aside",
    })
}

// ProvisionRequest describes incoming spec.
type ProvisionRequest struct {
    Engine    string `json:"engine"`
    Pattern   string `json:"pattern"`
    NodeCount int    `json:"nodeCount"`
}

// ProvisionResponse returns the project ID or queue token.
type ProvisionResponse struct {
    JobID string `json:"jobId"`
}

var prov = &provision.LocalProvisioner{} // swap for AWSProvisioner later

// Provision schedules the async provisioning job.
func Provision(c *gin.Context) {
    var req ProvisionRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 1. Generate a JobID (UUID, timestamp, etc.)
    jobID := "job-" + randString(8)

    // 2. Enqueue a background task (in real code, push to SQS)
    go func() {
        ctx := context.Background()
        spec := provision.Spec{Engine: req.Engine, Pattern: req.Pattern, NodeCount: req.NodeCount}
        result, err := prov.ProvisionCache(ctx, spec)
        // TODO: save `result` and `err` to DB keyed by jobID
        _ = result
        _ = err
    }()

    // 3. Return 202 Accepted with jobID
    c.JSON(http.StatusAccepted, ProvisionResponse{JobID: jobID})
}

// rudimentary random string for job IDs
func randString(n int) string {
    rand.Seed(time.Now().UnixNano())
    letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
    s := make([]rune, n)
    for i := range s {
        s[i] = letters[rand.Intn(len(letters))]
    }
    return string(s)
}
