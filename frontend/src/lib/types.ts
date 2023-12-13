import { z } from 'zod';

export const zUser = z.object({
	display: z.string(),
	username: z.string(),
	coins: z.number(),
	role: z.number(),
	joined_orgs: z.array(z.string().optional()).or(z.null())
});

export type User = z.infer<typeof zUser>;

export const zOrg = z.object({
	name: z.string(),
	tag: z.string(),
	is_unprotected: z.boolean(),
	adventures: z.array(z.string()).optional().or(z.null())
});

export type Org = z.infer<typeof zOrg>;

export const zAdventure = z.object({
	name: z.string(),
	tag: z.string(),
	description: z.string(),
	steps: z.array(z.string())
});

export type Adventure = z.infer<typeof zAdventure>;

export const Step = z.object({
	name: z.any(),
	tag: z.string(),
	description: z.any(),
	content: z.any(),
	category: z.string().default('step'),
	xp_award: z.number(),
	coins_award: z.number(),
	energy_cost: z.number()
});

// Functions

export const parseUser = (userData: Object): { user: User | null; error: boolean } => {
	const parsed = zUser.safeParse(userData);

	if (!parsed.success) {
		return { error: true, user: null };
	}
	return { user: parsed.data, error: false };
};

export const parseOrg = (orgObject: Object): { error: boolean; org: Org | null } => {
	const parsed = zOrg.safeParse(orgObject);

	if (!parsed.success) return { error: true, org: null };

	return { error: false, org: parsed.data };
};

export const parseAdventure = (
	adventureObject: Object
): { error: boolean; adventure: Adventure | null } => {
	const parsed = zAdventure.safeParse(adventureObject);

	if (!parsed.success) return { error: true, adventure: null };

	return { error: false, adventure: parsed.data };
};
