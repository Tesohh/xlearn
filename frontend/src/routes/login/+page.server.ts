import { authCookieName, backendUrl } from '$lib/const.js';
import { login, makeRequest } from '$lib/api/reqHandler.js';
import { fail, redirect } from '@sveltejs/kit';

export const load = async ({ locals }) => {
	if (locals.user) throw redirect(303, '/');
};

export const actions = {
	login: async ({ cookies, request }) => {
		console.log('LOGIN');

		const data = await request.formData();
		const username = data.get('username');
		const password = data.get('password');

		if (typeof username !== 'string' || typeof password !== 'string' || !username || !password) {
			return fail(400, { invalid: true });
		}

		const response = await login(username, password);

		if (response.error) {
			return fail(400, { error: true });
		}

		if (!response.cookie) return;

		cookies.set(authCookieName, response.cookie[0].replace('token=', ''), {
			httpOnly: true,
			maxAge: 60 * 60 * 24,
			secure: false, // TODO to change in production
			path: '/',
			sameSite: 'strict'
		});

		throw redirect(303, '/');
	}
};
