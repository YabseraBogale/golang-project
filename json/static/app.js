let point=document.querySelector(".point")


point.animate(
    [
        // keyframes
        { transform: "translateY(-20px)" },
        { transform: "translateY(20px)" },
        { transform: "translateX(20px)" },
        { transform: "translateX(-20px)" },
        { transform: "translateX(20px)" },
      ],
      {
        // timing options
        duration: 1000,
        iterations: Infinity,
      },
);

let low=document.getElementsByTagName("b");
if (low[3].innerText=="LOW"){
    low[3].style.color="green"
} else{
    low[3].style.color="red"
}