import { authCookieName } from '$lib/const.js';
import { redirect } from '@sveltejs/kit';

export const load = async ({ cookies, locals }) => {
	if (!locals.user) throw redirect(303, '/login');

	cookies.delete(authCookieName);

	throw redirect(303, '/login');
};
