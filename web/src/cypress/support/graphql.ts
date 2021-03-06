/* eslint-disable @typescript-eslint/no-explicit-any */
interface GraphQLResponse {
  data: any
  errors: [any]
}

// runs a graphql query returning the data response (after asserting no errors)
function graphql(query: string, variables?: any): Cypress.Chainable<any> {
  const url = '/api/graphql'
  if (!variables) variables = {}

  return cy.request('POST', url, { query, variables }).then((res) => {
    expect(res.status, 'status code').to.eq(200)

    let data: GraphQLResponse
    try {
      if (typeof res.body === 'string') data = JSON.parse(res.body)
      else data = res.body
    } catch (e) {
      console.error(res.body)
      throw e
    }
    if (data.errors && data.errors[0]) {
      // causes error message to be shown
      assert.isUndefined(data.errors[0].message)
    }
    expect(data).to.not.have.property('errors')

    return data.data
  })
}

Cypress.Commands.add('graphql', graphql)

export {}
