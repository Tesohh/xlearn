import { fail, redirect } from '@sveltejs/kit';

export const load = async ({ locals, url }) => {
	if (locals.user == null) throw redirect(303, '/login');

	console.log(url.searchParams.get('code'));
};

export const actions = {
	joinorg: async ({ locals }) => {
		return fail(400, { error: 'test' });
	}
};
