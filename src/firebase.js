import firebase from 'firebase/app';
import 'firebase/auth';

const firebaseConfig = {
  apiKey: "AIzaSyCFsR7mYKrS6KVeCg_1WiUe0GM8UDQBliY",
  authDomain: "academy-feedback-tool.firebaseapp.com",
  projectId: "academy-feedback-tool",
  storageBucket: "academy-feedback-tool.appspot.com",
  messagingSenderId: "376317359863",
  appId: "1:376317359863:web:4f27f65cc9f66f665945f2"
};

firebase.initializeApp(firebaseConfig);

export const auth = firebase.auth();
