import { authCookieName } from '$lib/const.js';
import { redirect } from '@sveltejs/kit';

export const load = async ({ cookies, locals }) => {
	if (!locals.user) throw redirect(303, '/login');

	cookies.delete(authCookieName, {
		secure: false // TODO change in production
	});

	throw redirect(303, '/login');
};
