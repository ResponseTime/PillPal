import React from "react";

export default function App() {
  const data = [
    { Day: "Monday", Time: "time" },
    { Day: "Tuesday", Time: "t76ime23" },
    { Day: "Wednesday", Time: "tim76e" },
    { Day: "Thursday", Time: "ti34me" },
    { Day: "Friday", Time: "ti4me" },
    { Day: "Saturday", Time: "ti7me" },
    { Day: "Sunday", Time: "t87ime" },
    { Day: "Monday", Time: "tim3e" },
    { Day: "Friday", Time: "tim00e" },
  ];
  const groupedData = {};

  data.forEach((item) => {
    const day = item.Day;
    if (!groupedData[day]) {
      groupedData[day] = [];
    }
    groupedData[day].push(item);
  });

  return (
    <>
      <div className="shower">
        {Object.keys(groupedData).map((elem, index) => {
          return (
            <div>
              <div className="day">{elem}</div>
              <div className="act">
                {groupedData[elem].map((item) => {
                  return <div>{item.Time}</div>;
                })}
              </div>
              ;
            </div>
          );
        })}
      </div>
    </>
  );
}
