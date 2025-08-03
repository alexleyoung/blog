
import { NextResponse } from 'next/server';
import admin from '@/lib/firebaseAdmin';
import { Post } from '@/lib/posts';

const postsCollection = admin.firestore().collection('posts');

export async function GET() {
    const snapshot = await postsCollection.orderBy('createdAt', 'desc').get();
    const posts = snapshot.docs.map(doc => {
        const data = doc.data();
        return {
            id: doc.id,
            ...data,
            createdAt: data.createdAt.toDate(),
            updatedAt: data.updatedAt.toDate(),
        };
    });
    return NextResponse.json(posts);
}

export async function POST(request: Request) {
    const body = await request.json();
    const now = new Date();
    const newPost: Omit<Post, 'id'> = {
        ...body,
        createdAt: now,
        updatedAt: now,
    };
    const docRef = await postsCollection.add(newPost);
    return NextResponse.json({ id: docRef.id, ...newPost }, { status: 201 });
}
