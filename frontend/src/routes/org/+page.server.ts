import { redirect } from '@sveltejs/kit';

export const load = async ({ locals }) => {
	if (!locals.user) throw redirect(303, '/login');

	throw redirect(303, '/');
};
