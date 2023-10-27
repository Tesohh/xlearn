import { env } from '$env/dynamic/private';
import { authCookieName } from '$lib/const.js';
import errorMessages from '$lib/errorMessages.js';
import { login, register } from '$lib/auth.js';
import { error, fail, redirect } from '@sveltejs/kit';

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
			return fail(400, { error: errorMessages.loginWrong });
		}

		if (password != confirmPassword) return fail(400, { error: errorMessages.registerError });
		if (password.length < 12) return fail(400, { error: errorMessages.registerError });

		const resp = await register(username, password);

		if (resp.error) return fail(400, { error: errorMessages.registerError });

		const respLogin = await login(username, password);

		if (respLogin.error) return fail(400, { error: errorMessages.somethingWentWrong });
		if (respLogin.cookie)
			cookies.set(authCookieName, respLogin.cookie, {
				httpOnly: true,
				maxAge: 60 * 60 * 24,
				secure: env.PRODUCTION ? true : false,
				path: '/',
				sameSite: 'strict'
			});
		else return fail(400, { error: errorMessages.somethingWentWrong });

		throw redirect(303, '/');
	}
};
