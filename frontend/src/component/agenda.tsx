import React from 'react'
type AgendaData = {
  title: string
  name: string
  time: number
}

export const Agenda = (): JSX.Element => {
  const [agenda, setAgenda] = React.useState<AgendaData[]>([
    {
      title: '自己紹介',
      name: 'うっしー',
      time: 3,
    },
    {
      title: '自己紹介',
      name: 'かわかみ',
      time: 3,
    },
    {
      title: 'うっしーからかわかみへ質問',
      name: 'うっしー',
      time: 3,
    },
    {
      title: '回答',
      name: 'かわかみ',
      time: 3,
    },
    {
      title: 'うっしーからかわかみへ質問',
      name: 'うっしー',
      time: 3,
    },
    {
      title: '回答',
      name: 'かわかみ',
      time: 3,
    },
    {
      title: 'おわりのあいさつ',
      name: 'うっしー',
      time: 3,
    },
    {
      title: 'おわりのあいさつ',
      name: 'かわかみ',
      time: 3,
    },
  ])

  return (
    <div className="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
      <div className="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th
                scope="col"
                className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                Title
              </th>
              <th
                scope="col"
                className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                スピーカー
              </th>
              <th
                scope="col"
                className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
              >
                時間
              </th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {agenda.map((a, index) => (
              <tr key={index}>
                <td className="px-6 py-4 whitespace-nowrap">
                  <div className="text-sm text-gray-900">{a.title}</div>
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <div className="text-sm text-gray-900">{a.name}</div>
                </td>
                <td className="px-6 py-4 whitespace-nowrap">
                  <div className="text-sm text-gray-500">{a.time}秒</div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}

export default Agenda
