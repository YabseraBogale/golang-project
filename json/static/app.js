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