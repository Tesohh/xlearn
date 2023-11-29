import { authCookieName } from '$lib/const.js';
import { getOrgByID, joinOrgByJoinCode } from '$lib/org.js';
import type { OrgType } from '$lib/types.js';
import { redirect } from '@sveltejs/kit';

export const load = async ({ locals, cookies }) => {
	if (!locals.user) throw redirect(303, '/login');

	if (locals.user?.joined_orgs == undefined) return { user: locals.user, org: null };

	const cookie = cookies.get(authCookieName);
	if (!cookie) return { user: null, org: null };

	const orgsData: OrgType[] = [];

	for (let i = 0; i < locals.user.joined_orgs.length; i++) {
		const orgID = locals.user.joined_orgs.at(i);

		if (!orgID) continue;

		let resp = await getOrgByID(orgID, cookie);
		if (resp.error) continue;
		if (resp.org == null) continue;
		orgsData.push(resp.org);
	}

	if (orgsData.length == 0) return { user: locals.user, org: null };
	return { user: locals.user, org: orgsData };
};

export const actions = {
	joinorg: async (event) => {
		if (!event.locals.user) return { failed: true };

		console.log('USER OK');

		const userCookie = event.cookies.get(authCookieName);

		if (!userCookie) return { failed: true };
		console.log('COOKIE FOUND');

		const data = await event.request.formData();
		console.log('FORMDATA FOUND');

		const joinCode = data.get('code')?.toString();

		if (joinCode == null) return { error: true };
		console.log('CODE FOUND');

		const resp = await joinOrgByJoinCode(joinCode, userCookie);

		if (resp.error) return { failed: true };

		return { success: true };
	}
};
