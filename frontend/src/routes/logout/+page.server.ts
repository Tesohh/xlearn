import { env } from '$env/dynamic/private';
import { authCookieName } from '$lib/const.js';
import { redirect } from '@sveltejs/kit';

export const load = async ({ cookies, locals }) => {
	if (!locals.user) redirect(303, '/login');

	/* @migration task: add path argument */ cookies.delete(authCookieName, {
		secure: env.PRODUCTION ? true : false,
		path: '/'
	});

	throw redirect(303, '/login');
};
