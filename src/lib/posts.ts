
import { firestore } from './firebaseAdmin';

export interface Post {
  id: string;
  title: string;
  content: string;
  createdAt: Date;
  updatedAt: Date;
}

const postsCollection = firestore.collection('posts');

export const createPost = async (post: Omit<Post, 'id' | 'createdAt' | 'updatedAt'>): Promise<Post> => {
  const now = new Date();
  const newPost = { ...post, createdAt: now, updatedAt: now };
  const docRef = await postsCollection.add(newPost);
  return { id: docRef.id, ...newPost };
};

export const getPost = async (id: string): Promise<Post | null> => {
  const doc = await postsCollection.doc(id).get();
  if (!doc.exists) {
    return null;
  }
  const data = doc.data();
  return { id: doc.id, ...data } as Post;
};

export const updatePost = async (id: string, post: Partial<Omit<Post, 'id' | 'createdAt'>>): Promise<Post> => {
  const now = new Date();
  const updatedPost = { ...post, updatedAt: now };
  await postsCollection.doc(id).update(updatedPost);
  const doc = await postsCollection.doc(id).get();
  const data = doc.data();
  return { id: doc.id, ...data } as Post;
};

export const deletePost = async (id: string): Promise<void> => {
  await postsCollection.doc(id).delete();
};

export const listPosts = async (): Promise<Post[]> => {
  const snapshot = await postsCollection.orderBy('createdAt', 'desc').get();
  return snapshot.docs.map(doc => {
    const data = doc.data();
    return { id: doc.id, ...data } as Post;
  });
};
