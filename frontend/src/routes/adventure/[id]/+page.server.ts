import { redirect } from '@sveltejs/kit';
import { backendUrl } from '$lib/const.js';

export const load = async ({ locals, params, fetch }) => {
	if (!locals.user) throw redirect(303, '/login');

	const result = await fetch(`${backendUrl}/api/org/ot/adventure/${params.id}`);

	if (!result.ok) throw redirect(303, '/error');
};
