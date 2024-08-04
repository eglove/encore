package blogMeta

import (
	"context"
	"encore.dev/storage/sqldb"
)

type BlogMetaInput struct {
	Slug          string   `json:"slug"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Tags          []string `json:"tags"`
	Author        string   `json:"author"`
	FeaturedImage string   `json:"featured_image"`
}

//encore:api auth method=POST path=/blog-meta
func createBlog(ctx context.Context, meta *BlogMetaInput) (*BlogMetaInput, error) {
	_, err := sqldb.Exec(ctx, `
		insert into BlogMeta (slug, title, description, tags, author, featuredImage)
		values ($1, $2, $3, $4, $5, $6)
	`, meta.Slug, meta.Title, meta.Description, meta.Tags, meta.Author, meta.FeaturedImage)

	if err != nil {
		return nil, err
	}

	return meta, nil
}

type GetBlogsResponse struct {
	Blogs []*BlogMetaInput `json:"blogs"`
}

// encore:api auth method=GET path=/blog-meta
func getBlogs(ctx context.Context) (*GetBlogsResponse, error) {
	rows, err := sqldb.Query(ctx, `
		select slug, title, description, tags, author, featuredImage
		from BlogMeta
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []*BlogMetaInput
	for rows.Next() {
		var blog BlogMetaInput
		err := rows.Scan(&blog.Slug, &blog.Title, &blog.Description, &blog.Tags, &blog.Author, &blog.FeaturedImage)
		if err != nil {
			return nil, err
		}
		blogs = append(blogs, &blog)
	}

	return &GetBlogsResponse{Blogs: blogs}, nil
}
