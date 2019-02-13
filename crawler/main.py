import os
from selenium import webdriver

PROJECT_ROOT = os.path.abspath(os.path.dirname(__file__))
DRIVER_BIN = os.path.join(PROJECT_ROOT, "driver/chromedriver")

base_url = 'http://www.avm.gen.tr'

browser = webdriver.Chrome(executable_path = DRIVER_BIN)
browser.get(base_url + '/avm/')

icerik = browser.find_elements_by_class_name('icerik')

for i in icerik:
    isim = i.find_element_by_class_name('isim').find_element_by_css_selector('a')
    map = i.find_element_by_class_name('map')

    isim_text = isim.get_attribute('title')
    map_text = map.text

    print(isim_text)
    print(map_text)