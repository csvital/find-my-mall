import os, time
from selenium import webdriver

PROJECT_ROOT = os.path.abspath(os.path.dirname(__file__))
DRIVER_BIN = os.path.join(PROJECT_ROOT, "driver/chromedriver")

base_url = 'https://www.avmgezgini.com/avmler/'

driver = webdriver.Chrome(executable_path = DRIVER_BIN)
driver.get(base_url)

boxes = driver.find_elements_by_class_name('avm_kutu')

for box in boxes:

    name = box.find_element_by_tag_name('p').text
    score = box.find_element_by_class_name('avm_puan').find_element_by_tag_name('b').text
    img_src = box.find_element_by_tag_name('img').get_attribute('src')
    detail_page = box.find_element_by_tag_name('a').get_attribute('href')

    print("Image link  : " + str(img_src))
    print("Detail page : " + str(detail_page))
    print("Avm name    : " + str(name))
    print("Avm score   : " + str(score))
    print()

    # Avm Detail Page
    driver.get(detail_page)
    magazalar_link = driver.find_element_by_xpath('//*[@id="EsitSolAlan"]/div[3]/a[1]').get_attribute('href')
    cafe_rest_link = driver.find_element_by_xpath('//*[@id="EsitSolAlan"]/div[3]/a[2]').get_attribute('href')

    # Avm List of ...
    driver.get(magazalar_link)
    dukkans = driver.find_elements_by_xpath('//*[@id="EsitSolAlan"]/div[5]/table/tbody/tr')
    for dukkan in dukkans:
        logo = dukkan.find_element_by_tag_name('img').get_attribute('src')
        magaza = dukkan.find_element_by_xpath('td[2]').text
        kat = dukkan.find_element_by_xpath('td[3]').text
        telefon = dukkan.find_element_by_xpath('td[4]').text
        #foto = dukkan.find_element_by_xpath('td[5]').text
        print("\t logo    : " + str(logo))
        print("\t magaza  : " + str(magaza))
        print("\t kat     : " + str(kat))
        print("\t telefon : " + str(telefon))
        #print("\t foto    : " + str(foto))
        print()

    driver.back()
    driver.get(cafe_rest_link)
    break

time.sleep(5)
driver.close()