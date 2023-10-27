import { authCookieName } from '$lib/const.js';
import { getOrgByID } from '$lib/org.js';
import type { OrgType } from '$lib/types.js';
import { redirect } from '@sveltejs/kit';

export const load = async ({ locals, cookies }) => {
	if (!locals.user) throw redirect(303, '/login');

	if (locals.user?.joined_orgs == undefined) return { user: locals.user, org: null };

	const cookie = cookies.get(authCookieName);
	if (!cookie) return { user: null, org: null };

	const orgsData: OrgType[] = [];

	for (let i = 0; i < locals.user.joined_orgs.length; i++) {
		let resp = await getOrgByID(locals.user.joined_orgs[i], cookie);
		if (resp.error) continue;
		if (resp.org == null) continue;
		orgsData.push(resp.org);
	}

	if (orgsData.length == 0) return { user: locals.user, org: null };
	return { user: locals.user, org: orgsData };
};
