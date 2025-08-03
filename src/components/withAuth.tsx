
'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { useAuthState } from 'react-firebase-hooks/auth';
import { auth } from '@/lib/firebase';

export default function withAuth(Component: React.ComponentType<any>) {
  return function AuthenticatedComponent(props: any) {
    const [user, loading] = useAuthState(auth);
    const router = useRouter();

    useEffect(() => {
      if (!loading && !user) {
        router.push('/login');
      }
    }, [user, loading, router]);

    if (loading) {
      return <div>Loading...</div>;
    }

    if (!user) {
      return null;
    }

    return <Component {...props} />;
  };
}
