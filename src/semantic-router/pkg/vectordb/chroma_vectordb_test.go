package vectordb

import "testing"

func TestChromaVectorDb(t *testing.T) {
	t.Run("ChromaVectorDb", func(t *testing.T) {
		options := ChromaVectorDbOptions{
			Endpoint:   "http://localhost:8000",
			Collection: "my_collection",
		}
		c, err := NewChromaVectorDb(options)
		if err != nil {
			t.Fatalf("Got error: %s", err)
		}
		t.Log("Vector DB instance created")
		res, err := c.Query("Hello")
		t.Log("Query complete")
		if err != nil {
			t.Fatalf("Got error: %s", err)
		}
		t.Log("Content:")
		for _, content := range res {
			t.Logf("%s \n", content)
		}
	})
}
