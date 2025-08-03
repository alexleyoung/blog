
import { initializeApp, getApps } from 'firebase/app';
import { getAuth } from 'firebase/auth';

const firebaseConfig = {
  apiKey: "AIzaSyA0PAYej3bGjiyb5GSP1_6ET0fqVWlvzuU",
  authDomain: "blog-47997.firebaseapp.com",
  projectId: "blog-47997",
  storageBucket: "blog-47997.firebasestorage.app",
  messagingSenderId: "914399541735",
  appId: "1:914399541735:web:63526d93662ee13557a542",
  measurementId: "G-QYMSKLN9L5"
};

if (!getApps().length) {
  initializeApp(firebaseConfig);
}

export const auth = getAuth();
