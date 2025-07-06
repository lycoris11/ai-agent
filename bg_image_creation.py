
import requests
from PIL import Image, ImageDraw, ImageFont
from io import BytesIO

def render_weather_component(day, icon_url, maxtemp, mintemp, condition, width=100, height=140, opacity=30):
    # Create a transparent image for the weather unit
    comp = Image.new('RGBA', (width, height), (0, 0, 0, 0))
    comp_draw = ImageDraw.Draw(comp)

    rect_radius = 10
    box_color = (113, 143, 166, opacity)
    comp_draw.rounded_rectangle([(5, 10), (width - 5, height - 10)], 
                                radius=rect_radius, fill=box_color)

    try:
        day_font = ImageFont.truetype("./assets/Roboto-Medium.ttf", 64)
        max_temp_font = ImageFont.truetype("./assets/Roboto-Medium.ttf", 52)
        min_temp_font = ImageFont.truetype("./assets/Roboto-Medium.ttf", 38)
        condition_font = ImageFont.truetype("./assets/Roboto-Medium.ttf", 42)
    except IOError:
        day_font = ImageFont.load_default()
        max_temp_font = ImageFont.load_default()
        min_temp_font = ImageFont.load_default()

    bbox = comp_draw.textbbox((0, 0), day, font=day_font)
    w, h = bbox[2] - bbox[0], bbox[3] - bbox[1]
    comp_draw.text(((width - w)//2, 30), day, fill=(225, 225, 225), font=day_font)

    icon = Image.open(BytesIO(requests.get(icon_url).content)).convert("RGBA")
    icon = icon.resize((128, 128), Image.LANCZOS)
    comp.paste(icon, ((width - 128)//2, 100), icon)
    
    bbox = comp_draw.textbbox((0, 0), maxtemp[:-1], font=max_temp_font)
    w2, h2 = bbox[2] - bbox[0], bbox[3] - bbox[1]
    comp_draw.text(((width - w2)//2, 255), maxtemp, fill=(225, 225, 225), font=max_temp_font, align="center")
    
    bbox = comp_draw.textbbox((0, 0), mintemp[:-1], font=min_temp_font)
    w3, h2 = bbox[2] - bbox[0], bbox[3] - bbox[1]
    comp_draw.text(((width - w3)//2, 310), mintemp, fill=(225, 225, 225), font=min_temp_font, align="center")
    
    bbox = comp_draw.textbbox((0, 0), condition, font=condition_font)
    w4, h3 = bbox[2] - bbox[0], bbox[3] - bbox[1]
    comp_draw.multiline_text(((width - w4)//2, 400), condition, fill=(225, 225, 225), font=condition_font, align="center")

    return comp

def create_bg_image(data):
    
    cell_width, cell_height = 274, 1080
    gap = 0
    n = len(data)
    total_width = cell_width*n + gap*(n-1)
    total_height = cell_height

    bg_image_url = "https://360chicago.com/wp-content/uploads/2024/05/Chicago-Weather-1.jpg"
    bg_resp = requests.get(bg_image_url)
    background = Image.open(BytesIO(bg_resp.content)).convert("RGBA")
    background = background.resize((total_width, total_height))

    result = background.copy()
    
    for i, item in enumerate(data):
        comp_img = render_weather_component(
            day=item["day"], 
            icon_url=item["icon_url"], 
            maxtemp=item["maxtemp"], 
            mintemp=item["mintemp"], 
            condition=item["condition"], 
            width=cell_width, 
            height=cell_height)
        x_pos = i*cell_width + i*gap
        result.paste(comp_img, (x_pos, 0), comp_img)

    result.save('./assets/temp_assets/weather_chicago.png')
    result.show()

