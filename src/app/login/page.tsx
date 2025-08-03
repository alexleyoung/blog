'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import { auth } from '@/lib/firebase';
import { createUserWithEmailAndPassword, signInWithEmailAndPassword } from 'firebase/auth';

export default function LoginPage() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [error, setError] = useState<string | null>(null);
  const [isSigningUp, setIsSigningUp] = useState(false);
  const router = useRouter();

  const handleAuth = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    if (isSigningUp) {
      if (password !== confirmPassword) {
        setError('Passwords do not match');
        return;
      }
      try {
        await createUserWithEmailAndPassword(auth, email, password);
        router.push('/admin');
      } catch (error: any) {
        setError(error.message);
      }
    } else {
      try {
        await signInWithEmailAndPassword(auth, email, password);
        router.push('/admin');
      } catch (error: any) {
        setError(error.message);
      }
    }
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-4xl font-bold mb-8">{isSigningUp ? 'Sign Up' : 'Login'}</h1>
      <form onSubmit={handleAuth}>
        <div className="mb-4">
          <label htmlFor="email" className="block text-lg font-bold mb-2">Email</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="w-full px-3 py-2 border rounded-lg"
            required
          />
        </div>
        <div className="mb-4">
          <label htmlFor="password" className="block text-lg font-bold mb-2">Password</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full px-3 py-2 border rounded-lg"
            required
          />
        </div>
        {isSigningUp && (
          <div className="mb-4">
            <label htmlFor="confirmPassword" className="block text-lg font-bold mb-2">Confirm Password</label>
            <input
              type="password"
              id="confirmPassword"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
              className="w-full px-3 py-2 border rounded-lg"
              required
            />
          </div>
        )}
        {error && <p className="text-red-500 mb-4">{error}</p>}
        <button type="submit" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
          {isSigningUp ? 'Sign Up' : 'Login'}
        </button>
      </form>
      <p className="mt-4">
        {isSigningUp ? 'Already have an account?' : 'Don\'t have an account?'}{' '}
        <button onClick={() => setIsSigningUp(!isSigningUp)} className="text-blue-500 hover:underline">
          {isSigningUp ? 'Login' : 'Sign Up'}
        </button>
      </p>
    </div>
  );
}
