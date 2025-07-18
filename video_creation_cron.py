import requests
import json
import datetime
import time
from bg_image_creation import create_bg_image
import os
from dotenv import load_dotenv

load_dotenv()

def format_condition(condition: str) -> str:
    return condition.replace(" ", "\n")

def format_temp(temp: float) -> str:
    return str(int(temp)) + "°"
    

def get_day_name(date_string: str) -> str:
    date_format = "%Y-%m-%d"
    date_object = datetime.datetime.strptime(date_string, date_format)
    day_of_week = date_object.weekday()
    day_name = ""
    
    
    if day_of_week == 0:
        day_name = "Mon"
    elif day_of_week == 1:
        day_name = "Tue"
    elif day_of_week == 2:
        day_name = "Wed"
    elif day_of_week == 3:
        day_name = "Thur"
    elif day_of_week == 4:
        day_name = "Fri"
    elif day_of_week == 5:
        day_name = "Sat"
    elif day_of_week == 6:
        day_name = "Sun"
    
    return day_name +" - " + f"{date_object.month}/{date_object.day}"

def get_data_for_image(weather_data) -> list:
    
    data = []
    
    for dateData in weather_data:
        dayInfo = {}
        dayInfo.update({'day': get_day_name(dateData['date'])})
        # dayInfo.update({'date': get_date(dateData['date'])})
        dayInfo.update({'icon_url': "https:"+dateData['day']['condition']['icon']})
        dayInfo.update({'maxtemp': format_temp(dateData['day']['maxtemp_f'])})
        dayInfo.update({'mintemp': format_temp(dateData['day']['mintemp_f'])})
        dayInfo.update({'condition': format_condition(dateData['day']['condition']['text'])})
        data.append(dayInfo)
    
    return data

def upload_image(file_path="./assets/temp_assets/weather_chicago.png") -> dict:
    
    url = "http://localhost:8080/video/backgroundImageUpload"

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
                    "type": "avatar",
                    "avatar_id": "Justo_Business_Front_public",
                    "scale": 0.51,
                    "offset": {
                        "x": 0.25,
                        "y": 0.25
                    },
                },
                "voice": {
                    "type": "text",
                    "voice_id": "9b6d89a2ac3f4a0eaa82f4d9ed9cabbf",
                    "input_text": script,
                    "speed": 1,
                    "elevenlabs_settings": {
                        "model": "eleven_multilingual_v2",
                        "similarity_boost": 0.75,
                        "stability" : 0.51,
                        "style": 0.0,
                    }
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
    print("Checking if video is generated.")
    url = f'http://localhost:8080//video/getStatus?video_id={video_id}'
    response = requests.get(url)
    res = json.loads(response.text)['data']['status']
    if res != 'completed' or res != 'failed':
        while(res == 'processing' or res == 'waiting' or res == 'pending'):
            time.sleep(120)
            print("Checking if video is generated.")
            response = requests.get(url)
            res = json.loads(response.text)['data']['status']
            
            if res == 'completed':
                break
            
            if res == 'failed':
                break
    return response.text

def download_video(url: str):
    
    with requests.get(url, stream=True) as r:
        r.raise_for_status()
        with open('video.mp4', 'wb') as f:
            for chunk in r.iter_content(chunk_size=8192):
                f.write(chunk)

def upload_video():
    
    today = datetime.date.today()
    
    payload = {
        "title": f'Chicago 3 Day Weather: {today}',
        "description": f'A 3 day weather forecast starting from {today}'
    }
    url = "http://localhost:8080/uploadVideo"
    
    with open("video.mp4", "rb") as f:
        files = {'file': ("video.mp4", f, "video/mp4")}
        response = requests.post(url, payload, files=files)
    
    print(response.text)
    return response.text

if __name__ == "__main__":
    #First get the 7 day weather data
    weatherData = json.loads(requests.get("http://localhost:8080/weather/7day/chicago").text)
    
    #Clean it up weather data. Use it to make a background image
    data = get_data_for_image(weatherData)
    create_bg_image(data)
    
    print("Uploading image to HeyGen")
    if os.getenv("ENV") == "dev":
        image_upload_response = json.loads(upload_image(file_path="./assets/weather_chicago.png"))
    else:
        image_upload_response = json.loads(upload_image(file_path="/home/ec2-user/ai-agent/assets/weather_chicago.png"))
    
    print("Getting the 7 day forecast script from gpt 4.1")
    script = get_7day_weather_script(weatherData)
    
    print("Generate the video on HeyGen using the background image and the script")
    video_response = json.loads(generate_video(image_url=image_upload_response["data"]["url"], script=script))
    
    video_id = video_response['data']['video_id']

    video_url = json.loads(check_if_video_is_generated(video_id))['data']['video_url']
    
    print("Downloading Video.")
    download_video(video_url)
    print("Uploading Video.")
    upload_video()
