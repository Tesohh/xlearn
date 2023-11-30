import { z } from 'zod';

export const User = z.object({
	display: z.string(),
	username: z.string(),
	coins: z.number(),
	role: z.number(),
	joined_orgs: z.array(z.string().optional()).or(z.null())
});

export type UserType = z.infer<typeof User>;

export const Org = z.object({
	name: z.string(),
	tag: z.string(),
	is_unprotected: z.boolean(),
	adventures: z.array(z.string()).optional().or(z.null())
});

export type OrgType = z.infer<typeof Org>;

// Functions

export const parseUser = (userData: Object): { user: UserType | null; error: boolean } => {
	const parsed = User.safeParse(userData);

	if (!parsed.success) {
		return { error: true, user: null };
	}
	return { user: parsed.data, error: false };
};

export const parseOrg = (orgObject: Object): { error: boolean; org: OrgType | null } => {
	const parsed = Org.safeParse(orgObject);

	if (!parsed.success) return { error: true, org: null };

	return { error: false, org: parsed.data };
};
