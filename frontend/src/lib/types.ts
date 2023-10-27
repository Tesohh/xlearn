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
