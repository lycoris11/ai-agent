#%%

import requests
import json
import datetime
from bg_image_creation import create_bg_image

def format_condition(condition: str) -> str:
    return condition.replace(" ", "\n")

def format_temp(temp: float) -> str:
    return str(int(temp)) + "°"
    

def get_day_name(date_string: str) -> str:
    date_format = "%Y-%m-%d"
    date_object = datetime.datetime.strptime(date_string, date_format)
    day_of_week = date_object.weekday()
    day_name = ""
    
    match day_of_week:
        case 0: day_name = "Mon"
        case 1: day_name = "Tue"
        case 2: day_name = "Wed"
        case 3: day_name = "Thur"
        case 4: day_name = "Fri"
        case 5: day_name = "Sat"
        case 6: day_name = "Sun"
    
    return day_name

def get_date_for_image() -> list:
    response = json.loads(requests.get("http://localhost:8080/weather/7day/chicago").text)
    
    data = []
    
    for dateData in response:
        dayInfo = {}
        dayInfo.update({'day': get_day_name(dateData['date'])})
        dayInfo.update({'icon_url': "https:"+dateData['day']['condition']['icon']})
        dayInfo.update({'maxtemp': format_temp(dateData['day']['maxtemp_f'])})
        dayInfo.update({'mintemp': format_temp(dateData['day']['mintemp_f'])})
        dayInfo.update({'condition': format_condition(dateData['day']['condition']['text'])})
        data.append(dayInfo)
    
    return data

if __name__ == "__main__":
    
    data = get_date_for_image()
    create_bg_image(data)
    #df = df.drop(['time_epoch', ], inplace=True)
    
    
    #print(json.dumps(response[0], indent=2))
# %%
