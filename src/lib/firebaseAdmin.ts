import * as admin from 'firebase-admin';

const serviceAccount = require('../../secrets/service_account_key.json');

if (!admin.apps.length) {
  admin.initializeApp({
    credential: admin.credential.cert(serviceAccount),
  });
}

export default admin;
