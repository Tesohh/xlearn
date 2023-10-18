import { env } from '$env/dynamic/private';
import { authCookieName, backendUrl } from '$lib/const.js';
import errorMessages from '$lib/errorMessages.js';
import { login } from '$lib/reqHandler.js';
import { fail, redirect } from '@sveltejs/kit';

export const load = async ({ locals }) => {
	if (locals.user) throw redirect(303, '/');
};

export const actions = {
	login: async ({ cookies, request }) => {
		const data = await request.formData();
		const username = data.get('username');
		const password = data.get('password');

		if (typeof username !== 'string' || typeof password !== 'string' || !username || !password) {
			return fail(400, { error: errorMessages.loginWrong });
		}

		const response = await login(username, password);

		if (response?.error) {
			return fail(400, { error: errorMessages.loginWrong });
		}

		if (response.cookie == null) return fail(400, { error: 'Error while logging in. Retry' });

		cookies.set(authCookieName, response.cookie, {
			httpOnly: true,
			maxAge: 60 * 60 * 24,
			secure: env.PRODUCTION ? true : false,
			path: '/',
			sameSite: 'strict'
		});

		throw redirect(303, '/');
	}
};
