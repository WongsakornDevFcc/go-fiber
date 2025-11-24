package middleware

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type TokenBucketConfig struct {
	Rate  int // tokens per second
	Burst int // max burst
	Redis *redis.Client
}

func TokenBucketMiddleware(cfg TokenBucketConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()
		user := c.Locals("jwt")

		var key string
		if user != nil {
			token := user.(*jwt.Token)
			claims := token.Claims.(jwt.MapClaims)
			email := claims["email"].(string)
			key = fmt.Sprintf("rate_limit:%s", email)
			log.Printf("Rate limiting for user: %s", email)
		} else {
			key = fmt.Sprintf("rate_limit:%s", c.IP())
			log.Printf("Using rate limit key: %s", key)
		}

		// Get current token info
		vals, err := cfg.Redis.HMGet(ctx, key, "tokens", "last_refill").Result()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		var tokens int
		var lastRefill time.Time

		if vals[0] == nil || vals[1] == nil {
			tokens = cfg.Burst
			lastRefill = time.Now()
		} else {
			tokens, _ = strconv.Atoi(vals[0].(string))
			ts, _ := strconv.ParseInt(vals[1].(string), 10, 64)
			lastRefill = time.Unix(ts, 0)
		}

		// Refill tokens
		elapsed := time.Since(lastRefill).Seconds()
		refill := int(elapsed) * cfg.Rate
		if refill > 0 {
			tokens += refill
			if tokens > cfg.Burst {
				tokens = cfg.Burst
			}
			lastRefill = time.Now()
		}

		if tokens <= 0 {
			retryAfter := time.Until(lastRefill.Add(time.Duration(cfg.Burst/cfg.Rate) * time.Second))
			c.Set("Retry-After", fmt.Sprintf("%.0f", retryAfter.Seconds()))
			return c.Status(fiber.StatusTooManyRequests).SendString("Rate limit exceeded")
		}

		// Consume a token
		tokens--
		cfg.Redis.HSet(ctx, key, "tokens", tokens)
		cfg.Redis.HSet(ctx, key, "last_refill", lastRefill.Unix())

		return c.Next()
	}
}
