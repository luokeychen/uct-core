package servers

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/pquerna/ffjson/ffjson"
	"strconv"
	"time"
	"uct/common/model"
)

const (
	jsonContentType     = "application/json; charset=utf-8"
	protobufContentType = "application/x-protobuf"

	ServingFromCache = "servingFromCache"
	ResponseKey      = "response"
	MetaKey          = "meta"

	contentTypeHeader   = "Content-Type"
	contentLengthHeader = "Content-Length"
)

func ProtobufWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if _, exists := c.Get(ServingFromCache); exists {
			return
		}

		if value, exists := c.Get(ResponseKey); exists {
			if response, ok := value.(model.Response); ok {
				// Write status header
				c.Writer.WriteHeader(int(*response.Meta.Code))

				// Serialize response
				b, err := response.Marshal()
				model.LogError(err)

				// Write Headers
				c.Header(contentLengthHeader, strconv.Itoa(len(b)))
				c.Header(contentTypeHeader, protobufContentType)

				// Write data and flush
				c.Writer.Write(b)
				c.Writer.Flush()
			}
		}
	}
}

func JsonWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if _, exists := c.Get(ServingFromCache); exists {
			return
		}

		if value, exists := c.Get(ResponseKey); exists {
			if response, ok := value.(model.Response); ok {
				// Write status header
				c.Writer.WriteHeader(int(*response.Meta.Code))
				// Serialize response
				b, err := ffjson.Marshal(response)
				model.LogError(err)

				// Write Headers
				c.Header(contentLengthHeader, strconv.Itoa(len(b)))
				c.Header(contentTypeHeader, jsonContentType)

				// Write data and flush
				c.Writer.Write(b)
				c.Writer.Flush()
			}
		}
	}
}

// Each handler must either set a meta or response
func ErrorWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		meta := model.Meta{}
		var metaExists bool
		var responseExists bool
		var value interface{}

		if value, metaExists = c.Get(MetaKey); metaExists {
			meta = value.(model.Meta)
		}

		if value, responseExists = c.Get(ResponseKey); responseExists {
			response, _ := value.(model.Response)
			code := int32(200)
			meta.Code = &code
			response.Meta = &meta
			c.Set(ResponseKey, response)
		} else {
			c.Set(ResponseKey, model.Response{Meta: &meta})
		}

		if !metaExists && !responseExists {
			c.Set(ServingFromCache, true)
		}
	}
}

func Ginrus(logger *logrus.Logger, timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		entry := logger.WithFields(logrus.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       path,
			"ip":         c.ClientIP(),
			"elapsed":    latency.Seconds() * 1e3,
			"latency":    latency,
			"user-agent": c.Request.UserAgent(),
			"time":       end.Format(timeFormat),
		})

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			entry.Error(c.Errors.String())
		} else {
			entry.Info()
		}
	}
}
