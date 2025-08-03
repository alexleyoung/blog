
import { NextResponse } from 'next/server';
import admin from '@/lib/firebaseAdmin';

const postsCollection = admin.firestore().collection('posts');

export async function GET(request: Request, { params }: { params: { id: string } }) {
    const doc = await postsCollection.doc(params.id).get();
    if (!doc.exists) {
        return NextResponse.json({ error: 'Post not found' }, { status: 404 });
    }
    const data = doc.data();
    const post = {
        id: doc.id,
        ...data,
        createdAt: data.createdAt.toDate(),
        updatedAt: data.updatedAt.toDate(),
    };
    return NextResponse.json(post);
}

export async function PUT(request: Request, { params }: { params: { id: string } }) {
    const body = await request.json();
    const now = new Date();
    const updatedPost = { ...body, updatedAt: now };
    await postsCollection.doc(params.id).update(updatedPost);
    const doc = await postsCollection.doc(params.id).get();
    const data = doc.data();
    const post = {
        id: doc.id,
        ...data,
        createdAt: data.createdAt.toDate(),
        updatedAt: data.updatedAt.toDate(),
    };
    return NextResponse.json(post);
}

export async function DELETE(request: Request, { params }: { params: { id: string } }) {
    await postsCollection.doc(params.id).delete();
    return new Response(null, { status: 204 });
}
