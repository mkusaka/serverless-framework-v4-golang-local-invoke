// handler.js
module.exports.hello = async (event) => {
    console.log('Received event:', JSON.stringify(event, null, 2));
  
    const response = {
      message: 'Hello from Lambda!',
      input: event,
    };
  
    console.log('Sending response:', JSON.stringify(response, null, 2));
    return response;
  };