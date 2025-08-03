
import { NextResponse } from 'next/server';
import { listPosts, createPost } from '@/lib/posts';

export async function GET() {
  const posts = await listPosts();
  return NextResponse.json(posts);
}

export async function POST(request: Request) {
  const body = await request.json();
  const newPost = await createPost(body);
  return NextResponse.json(newPost, { status: 201 });
}
