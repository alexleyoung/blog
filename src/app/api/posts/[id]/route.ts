
import { NextResponse } from 'next/server';
import { getPost, updatePost, deletePost } from '@/lib/posts';

export async function GET(request: Request, { params }: { params: { id: string } }) {
  const post = await getPost(params.id);
  if (!post) {
    return NextResponse.json({ error: 'Post not found' }, { status: 404 });
  }
  return NextResponse.json(post);
}

export async function PUT(request: Request, { params }: { params: { id: string } }) {
  const body = await request.json();
  const updatedPost = await updatePost(params.id, body);
  return NextResponse.json(updatedPost);
}

export async function DELETE(request: Request, { params }: { params: { id: string } }) {
  await deletePost(params.id);
  return new Response(null, { status: 204 });
}
