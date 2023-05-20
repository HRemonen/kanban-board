import React from 'react'
import { Link } from 'react-router-dom'
import { HiOutlineLightBulb } from 'react-icons/hi'

interface NewsCardProps {
  author: string
  title: string
  readTime: number
  info?: string
}

const NewsCard = ({ author, title, readTime, info }: NewsCardProps) => (
  <div className='space-y-2 rounded-xl border bg-white p-4 text-gray-800'>
    <div className='flex justify-between'>
      <div className='text-xs text-gray-400'>{author}</div>
      <div className='text-xs text-gray-400'>{readTime} min</div>
    </div>
    <Link to='/' className='font-bold hover:text-yellow-400 hover:underline'>
      {title}
    </Link>
    {info && (
      <div className='flex text-sm text-gray-600'>
        <HiOutlineLightBulb size={20} />
        <span className='ml-2'>{info}</span>
      </div>
    )}
  </div>
)

const NewsSection = () => (
  <section>
    <h2 className='mb-4 text-2xl font-bold'>What&lsquo;s new?</h2>
    <div className='space-y-4'>
      <NewsCard
        author='Pena Perala'
        title='New application launch!'
        readTime={4}
      />
      <NewsCard
        author='Kake Makela'
        title='Today is beerjantai!'
        readTime={2}
      />
      <NewsCard
        author='Uncle Bob'
        title='How to use Kanban in software development'
        readTime={10}
        info='Outstanding effort'
      />
    </div>
  </section>
)

export default NewsSection
