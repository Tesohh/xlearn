import { backendUrl } from '$lib/const.js';
import { login, makeRequest } from '$lib/reqHandler.js';
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

		cookies.set('jwt', response.cookie[0].replace('token=', ''));

		throw redirect(303, '/');
	}
};
