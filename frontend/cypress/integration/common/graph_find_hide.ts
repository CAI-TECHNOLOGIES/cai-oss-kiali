import { And, Then, When } from 'cypress-cucumber-preprocessor/steps';

function clearFindAndHide() {
  cy.get('#graph_hide').clear();
  cy.get('#graph_find').clear();
}

Then('user finds unhealthy workloads', () => {
  clearFindAndHide();
  cy.get('#graph_find').type('!healthy{enter}');
});

Then('user sees unhealthy workloads highlighted on the graph', () => {
  const expectedUnhealthyNodes = [
    {
      app: 'v-server',
      version: 'v1',
      namespace: 'alpha'
    },
    {
      app: 'w-server',
      version: 'v1',
      namespace: 'alpha'
    },
    {
      app: 'w-server',
      version: undefined, // Service does not have version
      namespace: 'alpha'
    }
  ];
  cy.waitForReact();
  cy.getReact('CytoscapeGraph')
    .should('have.length', '1')
    .getCurrentState()
    .then(state => {
      const unhealthyNodes = state.cy
        .nodes()
        .filter(node => node.classes().includes('find'))
        .map(node => ({ app: node.data('app'), version: node.data('version'), namespace: node.data('namespace') }));
      expect(unhealthyNodes).to.include.deep.members(expectedUnhealthyNodes);
    });
});

Then('user sees nothing highlighted on the graph', () => {
  cy.contains('Loading Graph').should('not.exist');

  cy.waitForReact();
  cy.getReact('CytoscapeGraph')
    .should('have.length', '1')
    .getCurrentState()
    .then(state => {
      expect(state.cy.nodes().filter(node => node.classes().includes('find')).length).to.equal(0);
    });
});

When('user hides unhealthy workloads', () => {
  clearFindAndHide();
  cy.get('#graph_hide').type('!healthy{enter}');
});

Then('user sees no unhealthy workloads on the graph', () => {
  cy.waitForReact();
  cy.getReact('CytoscapeGraph')
    .should('have.length', '1')
    .getCurrentState()
    .then(state => {
      const noUnhealthy = state.cy
        .nodes()
        // Unhealthy boxes are fine.
        .every(node => node.data('healthStatus') !== 'Failure' || node.data('nodeType') === 'box');
      expect(noUnhealthy).to.be.true;
    });
});

Then('user sees preset find options', () => {
  cy.getBySel('find-options-dropdown').click();
  cy.contains('Find: unhealthy nodes');
});

When('user selects the preset the find option {string}', (option: string) => {
  cy.get('#graph-find-presets').contains(option).click();
});

Then('user sees preset hide options', () => {
  cy.getBySel('hide-options-dropdown').click();
  cy.contains('Hide: healthy nodes');
});

When('user selects the preset hide option {string}', (option: string) => {
  cy.get('#graph-hide-presets').contains(option).click();
});

Then('user sees no healthy workloads on the graph', () => {
  cy.waitForReact();
  cy.getReact('CytoscapeGraph')
    .should('have.length', '1')
    .getCurrentState()
    .then(state => {
      const noHealthy = state.cy
        .nodes()
        .every(node => node.data('healthStatus') !== 'Healthy' || node.data('nodeType') === 'box');
      expect(noHealthy).to.be.true;
    });
});

When('user seeks help for find and hide', () => {
  cy.getBySel('graph-find-hide-help-button').click();
});

Then('user sees the help menu', () => {
  cy.getBySel('graph-find-hide-help').should('be.visible');
});

And('the help menu has info on {string}', (helpMenuItem: string) => {
  cy.get('#graph_find_help_tabs').contains(helpMenuItem).should('be.visible').click();
});
