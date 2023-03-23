import { baseUrl } from '../support/e2e'

describe('Front page', () => {
  it('is loading and responding to request', () => {
    cy.visit(baseUrl)
  })
})
