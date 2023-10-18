import { env } from '$env/dynamic/private';
import { authCookieName } from '$lib/const.js';
import { redirect } from '@sveltejs/kit';

export const load = async ({ cookies, locals }) => {
	if (!locals.user) throw redirect(303, '/login');

	cookies.delete(authCookieName, {
		secure: env.PRODUCTION ? true : false
	});

	throw redirect(303, '/login');
};
