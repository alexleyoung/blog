export interface Post {
  id: string;
  title: string;
  content: string;
  createdAt: Date;
  updatedAt: Date;
}

export const listPosts = async (): Promise<Post[]> => {
  const res = await fetch('/api/posts');
  if (!res.ok) {
    throw new Error('Failed to fetch posts');
  }
  const posts = await res.json();
  return posts.map((post: any) => ({
      ...post,
      createdAt: new Date(post.createdAt),
      updatedAt: new Date(post.updatedAt)
  }));
};

export const getPost = async (id: string): Promise<Post | null> => {
    const res = await fetch(`/api/posts/${id}`);
    if (!res.ok) {
        if (res.status === 404) {
            return null;
        }
        throw new Error('Failed to fetch post');
    }
    const post = await res.json();
    return {
        ...post,
        createdAt: new Date(post.createdAt),
        updatedAt: new Date(post.updatedAt)
    };
};

export const createPost = async (post: Omit<Post, 'id' | 'createdAt' | 'updatedAt'>): Promise<Post> => {
    const res = await fetch('/api/posts', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(post),
    });
    if (!res.ok) {
        throw new Error('Failed to create post');
    }
    const newPost = await res.json();
    return {
        ...newPost,
        createdAt: new Date(newPost.createdAt),
        updatedAt: new Date(newPost.updatedAt)
    };
};

export const updatePost = async (id: string, post: Partial<Omit<Post, 'id' | 'createdAt'>>): Promise<Post> => {
    const res = await fetch(`/api/posts/${id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(post),
    });
    if (!res.ok) {
        throw new Error('Failed to update post');
    }
    const updatedPost = await res.json();
    return {
        ...updatedPost,
        createdAt: new Date(updatedPost.createdAt),
        updatedAt: new Date(updatedPost.updatedAt)
    };
};

export const deletePost = async (id: string): Promise<void> => {
    const res = await fetch(`/api/posts/${id}`, {
        method: 'DELETE',
    });
    if (!res.ok) {
        throw new Error('Failed to delete post');
    }
};
