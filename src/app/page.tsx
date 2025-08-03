import Link from 'next/link';
import { listPosts, Post } from '@/lib/posts';

async function getPosts() {
  const posts = await listPosts();
  return posts;
}

export default async function Home() {
  const posts = await getPosts();

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center">
        <h1 className="text-4xl font-bold mb-8">Blog</h1>
        <Link href="/login" className="inline-block bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
          Login
        </Link>
      </div>
      <div className="grid gap-8">
        {posts.map((post: Post) => (
          <Link key={post.id} href={`/post/${post.id}`} className="block p-6 border rounded-lg hover:shadow-lg transition-shadow">
            <h2 className="text-2xl font-bold mb-2">{post.title}</h2>
            <p className="text-gray-600">{new Date(post.createdAt).toLocaleDateString()}</p>
          </Link>
        ))}
      </div>
    </div>
  );
}
