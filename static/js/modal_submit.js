
//Modal STuff
$(function(){
    $('#myForm').on('submit', function(e){
      e.preventDefault();
      console.log("Test test test");
      $.post('http://www.somewhere.com/path/to/post',
         $('#myForm').serialize(),
         function(data, status, xhr){
           // do something here with response;
         });
    });
});
