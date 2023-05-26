import { z } from 'zod'

export const NewBoardZod = z.object({
  name: z
    .string()
    .min(3, { message: 'Board name must be 3 or more characters long' })
    .max(255, { message: 'Board name must be 255 or less characters long' }),
  description: z.string().optional(),
})

export type NewBoard = z.infer<typeof NewBoardZod>
