import { env } from '$env/dynamic/private';
import { authCookieName } from '$lib/const.js';
import errorMessages from '$lib/errorMessages.js';
import { login, register } from '$lib/auth.js';
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
			return fail(400, { error: errorMessages.loginWrong, from: 'Login' });
		}

		const response = await login(username, password);

		if (response?.error) {
			return fail(400, { error: errorMessages.loginWrong, from: 'Login' });
		}

		if (response.cookie == null)
			return fail(400, { error: errorMessages.errorLoadingAuth, from: 'Login' });

		cookies.set(authCookieName, response.cookie, {
			httpOnly: true,
			maxAge: 60 * 60 * 24,
			secure: env.PRODUCTION ? true : false,
			path: '/',
			sameSite: 'strict'
		});

		throw redirect(303, '/');
	},

	signup: async ({ request, cookies }) => {
		const data = await request.formData();
		const username = data.get('username');
		const password = data.get('password');
		const confirmPassword = data.get('confirmpassword');

		if (
			typeof username !== 'string' ||
			typeof password !== 'string' ||
			typeof confirmPassword !== 'string' ||
			!username ||
			!password ||
			!confirmPassword
		) {
			return fail(400, { error: errorMessages.loginWrong, from: 'Signup' });
		}

		if (password != confirmPassword)
			return fail(400, { error: errorMessages.registerError, from: 'Signup' });
		if (password.length < 12)
			return fail(400, { error: errorMessages.registerError, from: 'Signup' });

		const resp = await register(username, password);

		if (resp.error) return fail(400, { error: errorMessages.registerError, from: 'Signup' });

		const respLogin = await login(username, password);

		if (respLogin.error) {
			return fail(400, { error: errorMessages.somethingWentWrong, from: 'Signup' });
		}
		if (respLogin.cookie) {
			cookies.set(authCookieName, respLogin.cookie, {
				httpOnly: true,
				maxAge: 60 * 60 * 24,
				secure: env.PRODUCTION ? true : false,
				path: '/',
				sameSite: 'strict'
			});
		} else return fail(400, { error: errorMessages.somethingWentWrong, from: 'Signup' });

		throw redirect(303, '/');
	}
};
