import { z } from 'zod'

export const NewBoardZod = z.object({
  name: z
    .string()
    .min(3, { message: 'Board name must be 3 or more characters long' })
    .max(255, { message: 'Board name must be 255 or less characters long' }),
  description: z.string().optional(),
})

export type NewBoard = z.infer<typeof NewBoardZod>

export const NewListZod = z.object({
  name: z
    .string()
    .min(1, { message: 'List name must be 1 or more characters long' })
    .max(255, { message: 'List name must be 255 or less characters long' }),
})

export type NewList = z.infer<typeof NewListZod>

export const NewCardZod = z.object({
  title: z
    .string()
    .min(3, { message: 'Card title must be 3 or more characters long' })
    .max(255, { message: 'Card title must be 255 or less characters long' }),
  description: z.string().optional(),
})

export type NewCard = z.infer<typeof NewCardZod>
