import { authCookieName } from '$lib/const';
import { getOrgByID } from '$lib/org';
import type { Org } from '$lib/types';

export const load = async ({ locals, cookies }) => {
	if (!locals.user) return { user: locals.user };

	if (locals.user?.joined_orgs == undefined) return { user: locals.user, org: null };

	const cookie = cookies.get(authCookieName);
	if (!cookie) return { user: null, org: null };

	const orgsData: Org[] = [];

	// Retrieving org data
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
