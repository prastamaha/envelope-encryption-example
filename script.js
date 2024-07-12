import http from 'k6/http';
import { sleep } from 'k6';
import { Trend, Rate } from 'k6/metrics';

let registerUserTrend = new Trend('register_user_trend');
let registerUserSuccessRate = new Rate('register_user_success_rate');
let getUserByUsernameTrend = new Trend('get_users_by_username_trend');
let getUserByUsernameSuccessRate = new Rate('get_users_by_username_success_rate');

export let options = {
    stages: [
        { duration: '1m', target: 10 }, // ramp up to 10 users
        { duration: '3m', target: 30 }, // ramp up to 30 users
        { duration: '1m', target: 0 }, // ramp down to 0 users
    ],
};

export default function () {
    // URL of the API server with encryption
    let urlRegisterUser = 'http://localhost:8000/users';
    let urlGetUserByUsername = 'http://localhost:8000/users/';

    // Generate random username
    let username = `user_${__VU}_${__ITER}`;

    // Payload for POST request
    let payload = {
        username: username,
        name: 'Super Cool User',
        gender: 'male',
        phone: '+628123456789',
        address: 'Jalan jalan keliling',
        consented: 'true',
    };

    // Headers for the request
    let headers = { 'Content-Type': 'application/x-www-form-urlencoded' };

    // Register user 
    let registerUserRespEnc = http.post(urlRegisterUser, payload, { headers });
    registerUserTrend.add(registerUserRespEnc.timings.duration);
    registerUserSuccessRate.add(registerUserRespEnc.status === 201);

    // Get user by username
    let getUserByUsernameRespEnc = http.get(urlGetUserByUsername + username);
    getUserByUsernameTrend.add(getUserByUsernameRespEnc.timings.duration);
    getUserByUsernameSuccessRate.add(getUserByUsernameRespEnc.status === 200);

    sleep(1);
}
