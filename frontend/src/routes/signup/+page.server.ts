import { authCookieName } from '$lib/const.js';
import { login, register } from '$lib/reqHandler.js';
import { fail, redirect } from '@sveltejs/kit';

export const load = async ({ locals }) => {
	if (locals.user) throw redirect(303, '/');
};

export const actions = {
	signup: async ({ request, cookies }) => {
		const data = await request.formData();
		const username = data.get('username');
		const password = data.get('password');
		const confirmPassword = data.get('confirmPassword');

		if (
			typeof username !== 'string' ||
			typeof password !== 'string' ||
			typeof confirmPassword !== 'string' ||
			!username ||
			!password ||
			!confirmPassword
		) {
			return fail(400, { invalid: true });
		}

		if (password != confirmPassword) return fail(400, { invalid: true });
		if (password.length < 12) return fail(400, { invalid: true });

		const resp = await register(username, password);

		if (resp.error) return fail(400, { invalid: true });

		const respLogin = await login(username, password);
		if (respLogin.error) return fail(400, { loginError: true });
		if (respLogin.cookie)
			cookies.set(authCookieName, respLogin.cookie, {
				httpOnly: true,
				maxAge: 60 * 60 * 24,
				secure: false, // TODO to change in production
				path: '/',
				sameSite: 'strict'
			});

		throw redirect(303, '/');
	}
};
