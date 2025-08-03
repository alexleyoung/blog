
'use client';

import Link from 'next/link';
import { useEffect, useState } from 'react';
import { listPosts, Post } from '@/lib/posts';
import withAuth from '@/components/withAuth';
import { auth } from '@/lib/firebase';
import { useRouter } from 'next/navigation';

function AdminPage() {
  const [posts, setPosts] = useState<Post[]>([]);
  const router = useRouter();

  const fetchPosts = async () => {
    const res = await fetch('/api/posts');
    const data = await res.json();
    setPosts(data);
  };

  useEffect(() => {
    fetchPosts();
  }, []);

  const handleDelete = async (id: string) => {
    if (window.confirm('Are you sure you want to delete this post?')) {
      await fetch(`/api/posts/${id}`, { method: 'DELETE' });
      fetchPosts();
    }
  };

  const handleSignOut = async () => {
    await auth.signOut();
    router.push('/login');
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-4xl font-bold">Admin</h1>
        <div>
          <Link href="/admin/new">
            <a className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded mr-4">
              New Post
            </a>
          </Link>
          <button onClick={handleSignOut} className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded">
            Sign Out
          </button>
        </div>
      </div>
      <div className="grid gap-4">
        {posts.map((post: Post) => (
          <div key={post.id} className="flex items-center justify-between p-4 border rounded-lg">
            <div>
              <h2 className="text-xl font-bold">{post.title}</h2>
              <p className="text-gray-600">{new Date(post.createdAt).toLocaleDateString()}</p>
            </div>
            <div>
              <Link href={`/admin/edit/${post.id}`}>
                <a className="text-blue-500 hover:underline mr-4">Edit</a>
              </Link>
              <button onClick={() => handleDelete(post.id)} className="text-red-500 hover:underline">
                Delete
              </button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default withAuth(AdminPage);
