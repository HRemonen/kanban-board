import React from 'react'
import { Droppable } from 'react-beautiful-dnd'

import { List } from '../../types'

const ListView = ({ list }: { list: List }) => {
  console.log('LISTVIEW', list)

  return (
    <Droppable droppableId={`list-${list.ID}-cards`}>
      {(provided) => (
        <ul
          data-cy={`list-${list.ID}-cards`}
          {...provided.droppableProps}
          ref={provided.innerRef}
        >
          {list.Cards.map((card) => (
            <div key={card.ID}>{card.Title}</div>
          ))}
        </ul>
      )}
    </Droppable>
  )
}

export default ListView
