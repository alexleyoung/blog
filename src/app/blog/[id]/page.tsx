
import { getPost, Post } from '@/lib/posts';
import { notFound } from 'next/navigation';

async function getPostData(id: string) {
  const post = await getPost(id);
  if (!post) {
    notFound();
  }
  return post;
}

export default async function PostPage({ params }: { params: { id: string } }) {
  const post = await getPostData(params.id);

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-4xl font-bold mb-4">{post.title}</h1>
      <p className="text-gray-600 mb-8">{new Date(post.createdAt).toLocaleDateString()}</p>
      <div className="prose lg:prose-xl">
        {post.content}
      </div>
    </div>
  );
}
