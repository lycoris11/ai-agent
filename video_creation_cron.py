import requests
import json
import datetime
import time
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

def get_data_for_image(weather_data) -> list:
    
    data = []
    
    for dateData in weather_data:
        dayInfo = {}
        dayInfo.update({'day': get_day_name(dateData['date'])})
        dayInfo.update({'icon_url': "https:"+dateData['day']['condition']['icon']})
        dayInfo.update({'maxtemp': format_temp(dateData['day']['maxtemp_f'])})
        dayInfo.update({'mintemp': format_temp(dateData['day']['mintemp_f'])})
        dayInfo.update({'condition': format_condition(dateData['day']['condition']['text'])})
        data.append(dayInfo)
    
    return data

def upload_image() -> dict:
    
    url = "http://localhost:8080/video/backgroundImageUpload"
    file_path = "./assets/temp_assets/weather_chicago.png"

    with open(file_path, 'rb') as f:
        files = {'file': f}
        response = requests.post(url, files=files)

    return response.text

def get_7day_weather_script(weatherData) -> str:
    url = "http://localhost:8080/ai/7day/weatherScript"
    payload = weatherData
    response = requests.post(url, json=payload)

    return response.text

def generate_video(image_url, script):
    url = "http://localhost:8080/video/generateVideo"
    payload = {
        "caption": False,
        "dimension": {
            "width": 1280,
            "height": 720
        },
        "video_inputs": [
            {
                "character": {
                    "type": "talking_photo",
                    "talking_photo_id": "aa80db9e21f5451aa3a31e7290c3cf9b",
                    "scale": 0.51,
                    "offset": {
                        "x": 0.25,
                        "y": 0.25
                    },
                },
                "voice": {
                    "type": "text",
                    "voice_id": "5f745b3db0db43739f31499f4f0aedd6",
                    "input_text": script,
                    "speed": 1.25
                },
                "background": {
                    "type": "image",
                    "url": image_url
                }
            }
        ]
    }
    response = requests.post(url, json=payload)
    #print('Status code:', response.status_code)
    #print('Response:', response.text)
    return response.text

def check_if_video_is_generated(video_id):
    url = f'http://localhost:8080//video/getStatus?video_id={video_id}'
    response = requests.get(url)
    res = json.loads(response.text)['data']['status']
    if res != 'completed' or res != 'failed':
        while(res == 'processing' or res == 'waiting' or res == 'pending'):
            time.sleep(120)
            response = requests.get(url)
            res = json.loads(response.text)['data']['status']
            
            if res == 'completed':
                break
            
            if res == 'failed':
                break
    return response.text

if __name__ == "__main__":
    #First get the 7 day weather data
    weatherData = json.loads(requests.get("http://localhost:8080/weather/7day/chicago").text)
    
    #Clean it up weather data. Use it to make a background image
    data = get_data_for_image(weatherData)
    create_bg_image(data)
    
    #Upload image to HeyGen
    response = json.loads(upload_image())
    
    #Get the 7 day forecast script from gpt 4.1
    script = get_7day_weather_script(weatherData)
    print(response["data"]["url"])
    print(script)
    
    #Generate the video on HeyGen using the background image and the script.
    vide_response = json.loads(generate_video(image_url=response["data"]["url"], script=script))
    video_id = vide_response['data']['video_id']

    video_url = json.loads(check_if_video_is_generated(video_id))['data']['video_url']
