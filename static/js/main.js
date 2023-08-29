$(document).ready(function(){
    console.log("hello world")
    $(".search-field").keyup(async function(){
         if($(".search-field").val().length==0){
            $("ul").empty()
         }
         else{
            objects=await fetch("http://localhost:8080/objects?search="+$(".search-field").val())
       let jsonObjArr=await objects.json()
       console.log(typeof(jsonObjArr))
        $("ul").empty()
        for(let i=0;i<jsonObjArr.length;i++){
            if(i==jsonObjArr.length-1){
                $("ul").append("<li class=`list-group-item`>"+jsonObjArr[i].Title+"</li><li class=`list-group-item` style='color:grey'>"+jsonObjArr[i].Description+"</li>")
                break;

            }
           
            $("ul").append("<li class=`list-group-item`>"+jsonObjArr[i].Title+"</li><li class=`list-group-item` style='color:grey'>"+jsonObjArr[i].Description+"</li>").append("<hr>")
        }
         }
    //     objects=await fetch("http://localhost:8080/objects?search="+$(".search-field").val())
    //    let jsonObjArr=await objects.json()
    //    console.log(typeof(jsonObjArr))
    //     $("ul").empty()
    //     for(let i=0;i<jsonObjArr.length;i++){
    //         console.log(jsonObjArr[i].Title)
           
    //         $("ul").append("<li class=`list-group-item`>"+jsonObjArr[i].Title+"</li><hr style='border:1px solid black'>")
    //     }
        
       
       
    })
});