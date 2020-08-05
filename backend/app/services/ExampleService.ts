export function listAll() {
  return [
    {
      id: 1,
      title: "Notion",
      link: "https://notion.so",
      description: "All in one tool to organize teams and ideas. Write, plan, collaborate, and get organized. ",
      tags: [ "organization", "planning", "collaboration", "writing", "calendar" ]
    },
    {
      id: 2,
      title: "json-server",
      link: "https://github.com/typicode/jsonserver",
      description: "Fake REST API based on a json schema. Useful for mocking and creating APIs for front-end devs to consume in coding challenges.", 
      tags: [ "api", "json", "schema", "node", "github", "rest" ]
    },
    { 
      id: 3,
      title: "fastify",
      link: "https://www.fastify.io/",
      description: "Extremely fast and simple, low-overhead web framework for NodeJS. Supports HTTP2.", 
      tags: [ "web", "framework", "node", "http2", "https", "localhost" ]
    }
  ];
}

export function store() {
  return ["Hi there!"];
}